package gameserver

import (
	agonesv1 "agones.dev/agones/pkg/apis/agones/v1"
)

type IngressRoutingMode string

const (
	IngressRoutingModeDomain IngressRoutingMode = "domain"
	IngressRoutingModePath   IngressRoutingMode = "path"

	OctopsAnnotationIngressMode    = "octops.io/gameserver-ingress-mode"
	OctopsAnnotationIngressDomain  = "octops.io/gameserver-ingress-domain"
	OctopsAnnotationIngressFQDN    = "octops.io/gameserver-ingress-fqdn"
	OctopsAnnotationTerminateTLS   = "octops.io/terminate-tls"
	OctopsAnnotationsTLSSecretName = "octops.io/tls-secret-name"
	OctopsAnnotationIssuerName     = "octops.io/issuer-tls-name"
	OctopsAnnotationCustomPrefix   = "octops-"
	AgonesAnnotationCustomPrefix   = "agones.dev/sdk-"

	CertManagerAnnotationIssuer = "cert-manager.io/issuer"
	AgonesGameServerNameLabel   = "agones.dev/gameserver"

	ErrGameServerAnnotationEmpty = "gameserver %s/%s has annotation %s but it is empty"
	ErrIngressRoutingModeEmpty   = "ingress routing mode %s requires the annotation %s to be set"
)

func (m IngressRoutingMode) String() string {
	return string(m)
}

func FromObject(obj interface{}) *agonesv1.GameServer {
	if gs, ok := obj.(*agonesv1.GameServer); ok {
		return gs
	}

	return &agonesv1.GameServer{}
}

func GetGameServerPort(gs *agonesv1.GameServer) agonesv1.GameServerStatusPort {
	if len(gs.Status.Ports) > 0 {
		return gs.Status.Ports[0]
	}

	return agonesv1.GameServerStatusPort{}
}

func GetGameServerContainerPort(gs *agonesv1.GameServer) int32 {
	if len(gs.Spec.Ports) > 0 {
		return gs.Spec.Ports[0].ContainerPort
	}

	return 0
}

func HasAnnotation(gs *agonesv1.GameServer, annotation string) (string, bool) {
	if value, ok := gs.Annotations[annotation]; ok {
		return value, true
	}

	return "", false
}

func IsReady(gs *agonesv1.GameServer) bool {
	if gs == nil {
		return false
	}

	return gs.Status.State == agonesv1.GameServerStateReady
}

func IsReadyOrAllocated(gs *agonesv1.GameServer) bool {
	if gs == nil {
		return false
	}

	return gs.Status.State == agonesv1.GameServerStateReady || gs.Status.State == agonesv1.GameServerStateAllocated
}

func GetIngressRoutingMode(gs *agonesv1.GameServer) IngressRoutingMode {
	if mode, ok := HasAnnotation(gs, OctopsAnnotationIngressMode); ok {
		return IngressRoutingMode(mode)
	}

	return IngressRoutingModeDomain
}

func GetTLSCertIssuer(gs *agonesv1.GameServer) string {
	if name, ok := HasAnnotation(gs, OctopsAnnotationIssuerName); ok {
		return name
	}

	return ""
}
