package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"0chain.net/blobbercore/challenge"
	"0chain.net/blobbercore/config"
	"0chain.net/blobbercore/datastore"
	"0chain.net/blobbercore/filestore"
	"0chain.net/blobbercore/handler"
	"0chain.net/blobbercore/readmarker"
	"0chain.net/blobbercore/writemarker"
	"0chain.net/core/build"
	"0chain.net/core/chain"
	"0chain.net/core/common"
	"0chain.net/core/encryption"
	"0chain.net/core/logging"
	. "0chain.net/core/logging"
	"0chain.net/core/node"
	"0chain.net/core/transaction"
	"0chain.net/core/util"

	"github.com/0chain/gosdk/zcncore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//var BLOBBER_REGISTERED_LOOKUP_KEY = datastore.ToKey("blobber_registration")

var startTime time.Time
var serverChain *chain.Chain
var filesDir *string
var metadataDB *string

func initHandlers(r *mux.Router) {
	r.HandleFunc("/", HomePageHandler)
	handler.SetupHandlers(r)
}

func SetupWorkerConfig() {
	config.Configuration.ContentRefWorkerFreq = viper.GetInt64("contentref_cleaner.frequency")
	config.Configuration.ContentRefWorkerTolerance = viper.GetInt64("contentref_cleaner.tolerance")

	config.Configuration.OpenConnectionWorkerFreq = viper.GetInt64("openconnection_cleaner.frequency")
	config.Configuration.OpenConnectionWorkerTolerance = viper.GetInt64("openconnection_cleaner.tolerance")

	config.Configuration.WMRedeemFreq = viper.GetInt64("writemarker_redeem.frequency")
	config.Configuration.WMRedeemNumWorkers = viper.GetInt("writemarker_redeem.num_workers")

	config.Configuration.RMRedeemFreq = viper.GetInt64("readmarker_redeem.frequency")
	config.Configuration.RMRedeemNumWorkers = viper.GetInt("readmarker_redeem.num_workers")

	config.Configuration.ChallengeResolveFreq = viper.GetInt64("challenge_response.frequency")
	config.Configuration.ChallengeResolveNumWorkers = viper.GetInt("challenge_response.num_workers")
	config.Configuration.ChallengeMaxRetires = viper.GetInt("challenge_response.max_retries")

	config.Configuration.ColdStorageMinimumFileSize = viper.GetInt64("cold_storage.min_file_size")
	config.Configuration.ColdStorageTimeLimitInHours = viper.GetInt64("cold_storage.file_time_limit_in_hours")
	config.Configuration.ColdStorageJobQueryLimit = viper.GetInt64("cold_storage.job_query_limit")
	config.Configuration.MaxCapacityPercentage = viper.GetFloat64("cold_storage.max_capacity_percentage")

	config.Configuration.MinioStart = viper.GetBool("minio.start")
	config.Configuration.MinioWorkerFreq = viper.GetInt64("minio.worker_frequency")
	config.Configuration.MinioNumWorkers = viper.GetInt("minio.num_workers")
	config.Configuration.MinioUseSSL = viper.GetBool("minio.use_ssl")

	config.Configuration.Capacity = viper.GetInt64("capacity")

	config.Configuration.DBHost = viper.GetString("db.host")
	config.Configuration.DBName = viper.GetString("db.name")
	config.Configuration.DBPort = viper.GetString("db.port")
	config.Configuration.DBUserName = viper.GetString("db.user")
	config.Configuration.DBPassword = viper.GetString("db.password")

	config.Configuration.Capacity = viper.GetInt64("capacity")
	config.Configuration.ReadPrice = viper.GetFloat64("read_price")
	config.Configuration.WritePrice = viper.GetFloat64("write_price")
	config.Configuration.MinLockDemand = viper.GetFloat64("min_lock_demand")
	config.Configuration.MaxOfferDuration = viper.GetDuration("max_offer_duration")
	config.Configuration.ChallengeCompletionTime = viper.GetDuration("challenge_completion_time")

	config.Configuration.FaucetWorkerFreqInMinutes = viper.GetInt64("faucet.worker_frequency")
	config.Configuration.FaucetMinimumBalance = viper.GetFloat64("faucet.minimum_balance")

	config.Configuration.ReadLockTimeout = int64(
		viper.GetDuration("read_lock_timeout") / time.Second,
	)
	config.Configuration.WriteLockTimeout = int64(
		viper.GetDuration("write_lock_timeout") / time.Second,
	)
}

func SetupWorkers() {
	handler.SetupWorkers(common.GetRootContext())
	challenge.SetupWorkers(common.GetRootContext())
	readmarker.SetupWorkers(common.GetRootContext())
	writemarker.SetupWorkers(common.GetRootContext())
	// stats.StartEventDispatcher(2)
}

var fsStore filestore.FileStore

func initEntities() {
	// badgerdbstore.SetupStorageProvider(*badgerDir)
	fsStore = filestore.SetupFSStore(*filesDir + "/files")
	// blobber.SetupObjectStorageHandler(fsStore, badgerdbstore.GetStorageProvider())

	// allocation.SetupAllocationChangeCollectorEntity(badgerdbstore.GetStorageProvider())
	// allocation.SetupAllocationEntity(badgerdbstore.GetStorageProvider())
	// allocation.SetupDeleteTokenEntity(badgerdbstore.GetStorageProvider())
	// reference.SetupFileRefEntity(badgerdbstore.GetStorageProvider())
	// reference.SetupRefEntity(badgerdbstore.GetStorageProvider())
	// reference.SetupContentReferenceEntity(badgerdbstore.GetStorageProvider())
	// writemarker.SetupEntity(badgerdbstore.GetStorageProvider())
	// readmarker.SetupEntity(badgerdbstore.GetStorageProvider())
	// challenge.SetupEntity(badgerdbstore.GetStorageProvider())
	// stats.SetupStatsEntity(badgerdbstore.GetStorageProvider())
}

func initServer() {

}

func checkForDBConnection() {
	retries := 0
	var err error
	for retries < 600 {
		err = datastore.GetStore().Open()
		if err != nil {
			time.Sleep(1 * time.Second)
			retries++
			continue
		}
		break
	}

	if err != nil {
		Logger.Error("Error in opening the database. Shutting the server down")
		panic(err)
	}
}

func processBlockChainConfig(nodesFileName string) {
	nodeConfig := viper.New()
	nodeConfig.AddConfigPath("./keysconfig")
	nodeConfig.AddConfigPath("./config")
	nodeConfig.SetConfigName(nodesFileName)

	err := nodeConfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	config := nodeConfig.Get("miners")
	if miners, ok := config.([]interface{}); ok {
		serverChain.Miners.AddNodes(miners)
	}
	config = nodeConfig.Get("sharders")
	if sharders, ok := config.([]interface{}); ok {
		serverChain.Sharders.AddNodes(sharders)
	}
}

func processMinioConfig(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)
	more := scanner.Scan()
	if more == false {
		return common.NewError("process_minio_config_failed", "Unable to read minio config from minio config file")
	}

	filestore.MinioConfig.StorageServiceURL = scanner.Text()
	more = scanner.Scan()
	if more == false {
		return common.NewError("process_minio_config_failed", "Unable to read minio config from minio config file")
	}

	filestore.MinioConfig.AccessKeyID = scanner.Text()
	more = scanner.Scan()
	if more == false {
		return common.NewError("process_minio_config_failed", "Unable to read minio config from minio config file")
	}

	filestore.MinioConfig.SecretAccessKey = scanner.Text()
	more = scanner.Scan()
	if more == false {
		return common.NewError("process_minio_config_failed", "Unable to read minio config from minio config file")
	}

	filestore.MinioConfig.BucketName = scanner.Text()
	more = scanner.Scan()
	if more == false {
		return common.NewError("process_minio_config_failed", "Unable to read minio config from minio config file")
	}

	filestore.MinioConfig.BucketLocation = scanner.Text()
	return nil
}

func isValidOrigin(origin string) bool {
	var url, err = url.Parse(origin)
	if err != nil {
		return false
	}
	var host = url.Hostname()
	if host == "localhost" {
		return true
	}
	if host == "0chain.net" ||
		strings.HasSuffix(host, ".0chain.net") ||
		strings.HasSuffix(host, ".alphanet-0chain.net") ||
		strings.HasSuffix(host, ".testnet-0chain.net") ||
		strings.HasSuffix(host, ".devnet-0chain.net") ||
		strings.HasSuffix(host, ".mainnet-0chain.net") {
		return true
	}
	return false
}

func main() {
	deploymentMode := flag.Int("deployment_mode", 2, "deployment_mode")
	nodesFile := flag.String("nodes_file", "", "nodes_file")
	keysFile := flag.String("keys_file", "", "keys_file")
	minioFile := flag.String("minio_file", "", "minio_file")
	filesDir = flag.String("files_dir", "", "files_dir")
	metadataDB = flag.String("db_dir", "", "db_dir")
	logDir := flag.String("log_dir", "", "log_dir")
	portString := flag.String("port", "", "port")
	hostname := flag.String("hostname", "", "hostname")

	flag.Parse()

	config.SetupDefaultConfig()
	config.SetupConfig()

	config.Configuration.DeploymentMode = byte(*deploymentMode)

	if config.Development() {
		logging.InitLogging("development", *logDir, "0chainBlobber.log")
	} else {
		logging.InitLogging("production", *logDir, "0chainBlobber.log")
	}
	config.Configuration.ChainID = viper.GetString("server_chain.id")
	config.Configuration.SignatureScheme = viper.GetString("server_chain.signature_scheme")
	SetupWorkerConfig()

	if *filesDir == "" {
		panic("Please specify --files_dir absolute folder name option where uploaded files can be stored")
	}

	if *metadataDB == "" {
		panic("Please specify --db_dir absolute folder name option where meta data db can be stored")
	}

	if *hostname == "" {
		panic("Please specify --hostname which is the public hostname")
	}

	if *portString == "" {
		panic("Please specify --port which is the port on which requests are accepted")
	}

	reader, err := os.Open(*keysFile)
	if err != nil {
		panic(err)
	}

	publicKey, privateKey, _, _ := encryption.ReadKeys(reader)
	reader.Close()

	reader, err = os.Open(*minioFile)
	if err != nil {
		panic(err)
	}

	err = processMinioConfig(reader)
	if err != nil {
		panic(err)
	}
	reader.Close()

	node.Self.SetKeys(publicKey, privateKey)

	port, err := strconv.Atoi(*portString) //fmt.Sprintf(":%v", port) // node.Self.Port
	if err != nil {
		Logger.Panic("Port specified is not Int " + *portString)
		return
	}

	node.Self.SetHostURL(*hostname, port)
	Logger.Info(" Base URL" + node.Self.GetURLBase())

	config.SetServerChainID(config.Configuration.ChainID)

	common.SetupRootContext(node.GetNodeContext())
	//ctx := common.GetRootContext()
	serverChain = chain.NewChainFromConfig()

	if *nodesFile == "" {
		panic("Please specify --nodes_file file.txt option with a file.txt containing nodes including self")
	}

	if strings.HasSuffix(*nodesFile, "txt") {
		reader, err = os.Open(*nodesFile)
		if err != nil {
			log.Fatalf("%v", err)
		}

		node.ReadNodes(reader, serverChain.Miners, serverChain.Sharders, serverChain.Blobbers)
		reader.Close()
	} else { //assumption it has yaml extension
		processBlockChainConfig(*nodesFile)
	}

	if node.Self.ID == "" {
		Logger.Panic("node definition for self node doesn't exist")
	} else {
		Logger.Info("self identity", zap.Any("id", node.Self.Node.GetKey()))
	}

	//address := publicIP + ":" + portString
	address := ":" + *portString

	chain.SetServerChain(serverChain)

	serverChain.Miners.ComputeProperties()
	serverChain.Sharders.ComputeProperties()
	serverChain.Blobbers.ComputeProperties()

	checkForDBConnection()

	// Initializa after serverchain is setup.
	initEntities()
	//miner.GetMinerChain().SetupGenesisBlock(viper.GetString("server_chain.genesis_block.id"))
	SetupBlobberOnBC(*logDir)
	mode := "main net"
	if config.Development() {
		mode = "development"
	} else if config.TestNet() {
		mode = "test net"
	}
	Logger.Info("Starting blobber", zap.Int("available_cpus", runtime.NumCPU()), zap.String("port", *portString), zap.String("chain_id", config.GetServerChainID()), zap.String("mode", mode))

	var server *http.Server

	// setup CORS
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{
		"X-Requested-With", "X-App-Client-ID",
		"X-App-Client-Key", "Content-Type",
	})
	originsOk := handlers.AllowedOriginValidator(isValidOrigin)
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT",
		"DELETE", "OPTIONS"})

	initHandlers(r)
	initServer()

	rHandler := handlers.CORS(originsOk, headersOk, methodsOk)(r)
	if config.Development() {
		// No WriteTimeout setup to enable pprof
		server = &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 30 * time.Second,
			MaxHeaderBytes:    1 << 20,
			Handler:           rHandler,
		}
	} else {
		server = &http.Server{
			Addr:              address,
			ReadHeaderTimeout: 30 * time.Second,
			WriteTimeout:      30 * time.Second,
			IdleTimeout:       30 * time.Second,
			MaxHeaderBytes:    1 << 20,
			Handler:           rHandler,
		}
	}
	common.HandleShutdown(server)
	handler.HandleShutdown(common.GetRootContext())

	Logger.Info("Ready to listen to the requests")
	startTime = time.Now().UTC()
	log.Fatal(server.ListenAndServe())
}

func RegisterBlobber() {

	registrationRetries := 0
	//ctx := badgerdbstore.GetStorageProvider().WithConnection(common.GetRootContext())
	for registrationRetries < 10 {
		txnHash, err := handler.RegisterBlobber(common.GetRootContext())
		time.Sleep(transaction.SLEEP_FOR_TXN_CONFIRMATION * time.Second)
		txnVerified := false
		verifyRetries := 0
		for verifyRetries < util.MAX_RETRIES {
			time.Sleep(transaction.SLEEP_FOR_TXN_CONFIRMATION * time.Second)
			t, err := transaction.VerifyTransaction(txnHash, chain.GetServerChain())
			if err == nil {
				txnVerified = true
				Logger.Info("Transaction for adding blobber accepted and verified", zap.String("txn_hash", t.Hash), zap.Any("txn_output", t.TransactionOutput))
				//badgerdbstore.GetStorageProvider().WriteBytes(ctx, BLOBBER_REGISTERED_LOOKUP_KEY, []byte(txnHash))
				//badgerdbstore.GetStorageProvider().Commit(ctx)
				SetupWorkers()
				go BlobberHealthCheck()
				return
			}
			verifyRetries++
		}

		if !txnVerified {
			Logger.Error("Add blobber transaction could not be verified", zap.Any("err", err), zap.String("txn.Hash", txnHash))
		}
	}
}

func BlobberHealthCheck() {
	const HEALTH_CHECK_TIMER = 60 * 15 // 15 Minutes
	for {
		txnHash, err := handler.BlobberHealthCheck(common.GetRootContext())
		if err != nil && err == handler.ErrBlobberHasRemoved {
			time.Sleep(HEALTH_CHECK_TIMER * time.Second)
			continue
		}
		time.Sleep(transaction.SLEEP_FOR_TXN_CONFIRMATION * time.Second)
		txnVerified := false
		verifyRetries := 0
		for verifyRetries < util.MAX_RETRIES {
			time.Sleep(transaction.SLEEP_FOR_TXN_CONFIRMATION * time.Second)
			t, err := transaction.VerifyTransaction(txnHash, chain.GetServerChain())
			if err == nil {
				txnVerified = true
				Logger.Info("Transaction for blobber health check verified", zap.String("txn_hash", t.Hash), zap.Any("txn_output", t.TransactionOutput))
				break
			}
			verifyRetries++
		}

		if !txnVerified {
			Logger.Error("Blobber health check transaction could not be verified", zap.Any("err", err), zap.String("txn.Hash", txnHash))
		}
		time.Sleep(HEALTH_CHECK_TIMER * time.Second)
	}
}

func SetupBlobberOnBC(logDir string) {
	var logName = logDir + "/0chainBlobber.log"
	zcncore.SetLogFile(logName, false)
	zcncore.SetLogLevel(3)
	zcncore.InitZCNSDK(serverChain.Miners.GetBaseURLsArray(), serverChain.Sharders.GetBaseURLsArray(), config.Configuration.SignatureScheme)
	zcncore.SetWalletInfo(node.Self.GetWalletString(), false)
	//txnHash, err := badgerdbstore.GetStorageProvider().ReadBytes(common.GetRootContext(), BLOBBER_REGISTERED_LOOKUP_KEY)
	//if err != nil {
	// Now register blobber to chain
	if config.Development() {
		CheckForFunds()
	}
	go RegisterBlobber()
	//}
	//Logger.Info("Blobber already registered", zap.Any("blobberTxn", string(txnHash)))
}

/*HomePageHandler - provides basic info when accessing the home page of the server */
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	mc := chain.GetServerChain()
	fmt.Fprintf(w, "<div>Running since %v ...\n", startTime)
	fmt.Fprintf(w, "<div>Working on the chain: %v</div>\n", mc.ID)
	fmt.Fprintf(w, "<div>I am a blobber with <ul><li>id:%v</li><li>public_key:%v</li><li>build_tag:%v</li></ul></div>\n", node.Self.GetKey(), node.Self.PublicKey, build.BuildTag)
	serverChain.Miners.Print(w)
	serverChain.Sharders.Print(w)
	serverChain.Blobbers.Print(w)
}

func CheckForFunds() {
	balance, err := handler.CheckBalance()
	if err != nil {
		Logger.Error("Failed to check for funds", zap.Error(err))
		panic("Unable to get balance")
	}
	for balance < config.Configuration.FaucetMinimumBalance {
		Logger.Info("Doesn't have minimum balance required, Calling faucet")
		err = handler.CallFaucet()
		if err != nil {
			Logger.Error("Failed to call faucet", zap.Error(err))
			continue
		}
		balance, err = handler.CheckBalance()
		if err != nil {
			Logger.Error("Failed to check for funds", zap.Error(err))
			panic("Unable to get balance")
		}
		Logger.Info("Faucet successfully called", zap.Any("current_balance", balance))
	}
}
