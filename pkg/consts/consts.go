package consts
import (
    "os"
)
type CustomError string

func (e CustomError) Error() string { return string(e) }


var (
    IngressClassPrivate = getEnv("INGRESS_CLASS_PRIVATE", "private")
	IngressClassPublic = getEnv("INGRESS_CLASS_PUBLIC", "public")
    IngressClassInterDc = getEnv("INGRESS_CLASS_INTER_DC", "inter-dc")
	TLSSecretNS = getEnv("TLS_SECRET_NS", "openshift-ingress")
	TLSSecretName = getEnv("TLS_SECRET_NAME","letsencrypt")
)

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
const (
	LabelKeyRouterName = "router"

	haproxyAnnotationPrefix = "haproxy.router.openshift.io/"

	AnnotRateLimit         = haproxyAnnotationPrefix + "rate-limit-connections"
	AnnotRateLimitHttpRate = haproxyAnnotationPrefix + "rate-limit-connections.rate-http"
	AnnotBalance           = haproxyAnnotationPrefix + "balance"
	AnnotTimeout           = haproxyAnnotationPrefix + "timeout"
	AnnotIPWhitelist       = haproxyAnnotationPrefix + "ip_whitelist"

	AnnotationKeyPrefix               = "snappcloud.io/"
	AnnotationKeyReconciliationPaused = AnnotationKeyPrefix + "paused"
	AnnotationKeyHttp1Enforced        = AnnotationKeyPrefix + "force-h1"


	GlobalTLSSecretName = TLSSecretNS + "/" + TLSSecretName

	RateLimitUnitMinute = "minute"

	RouteShardLabel = "router"

	StrategyCookie               = "Cookie"
	StrategyRandom               = "Random"
	StrategyRoundRobin           = "RoundRobin"
	StrategyWeightedLeastRequest = "WeightedLeastRequest"
	StrategyRequestHash          = "RequestHash"
	StrategyDefault              = StrategyRoundRobin

	TLSSecretGenerateName = "managed-tls-secret-"

	NotFoundError = CustomError("not found")
)
