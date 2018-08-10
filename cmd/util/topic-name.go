package util

import (
	"strings"
	"log"
	"fmt"
	"net/url"
)

type TopicName struct {
	completeTopicName string

	domain           string
	tenant           string
	cluster          string
	namespacePortion string
	localName        string
}

const Persistent = "persistent"
const NonPersistent = "non-persistent"

const PublicTenant = "public"
const DefaultNamespace = "default"
const PartitionedTopicSuffix = "-partition-"

func TopicNameParse(completeTopicName string) *TopicName {
	topicName := TopicName{}

	// The topic name can be in two different forms, one is fully qualified topic name,
	// the other one is short topic name
	if !strings.Contains(completeTopicName, "://") {
		// The short topic name can be:
		// - <topic>
		// - <property>/<namespace>/<topic>
		parts := strings.Split(completeTopicName, "/")
		if len(parts) == 3 {
			completeTopicName = Persistent + "://" + completeTopicName
		} else if len(parts) == 1 {
			completeTopicName = Persistent + "://" + PublicTenant + "/" + DefaultNamespace + "/" + parts[0]
		} else {
			log.Fatal(
				"Invalid short topic name '" + completeTopicName + "', it should be in the format of <tenant>/<namespace>/<topic> or <topic>")
		}
	}

	topicName.completeTopicName = completeTopicName

	// The fully qualified topic name can be in two different forms:
	// new:    persistent://tenant/namespace/topic
	// legacy: persistent://tenant/cluster/namespace/topic

	parts := strings.SplitN(completeTopicName, "://", 2)
	topicName.domain = parts[0]
	if topicName.domain != Persistent && topicName.domain != NonPersistent {
		log.Fatal("Invalid topic domain: ", topicName.domain)
	}

	rest := parts[1]

	// The rest of the name can be in different forms:
	// new:    tenant/namespace/<localName>
	// legacy: tenant/cluster/namespace/<localName>
	// Examples of localName:
	// 1. some/name/xyz//
	// 2. /xyz-123/feeder-2
	parts = strings.SplitN(rest, "/", 4)
	if len(parts) == 3 {
		// New topic name without cluster name
		topicName.tenant = parts[0]
		topicName.cluster = ""
		topicName.namespacePortion = parts[1]
		topicName.localName = parts[2]
	} else if len(parts) == 4 {
		// Legacy topic name that includes cluster name
		topicName.tenant = parts[0]
		topicName.cluster = parts[1]
		topicName.namespacePortion = parts[2]
		topicName.localName = parts[3]
	} else {
		log.Fatal("Invalid topic name: " + completeTopicName);
	}

	return &topicName
}

func (topicName *TopicName) isV2() bool {
	return topicName.cluster == ""
}

func (topicName *TopicName) RestPath() string {
	if topicName.isV2() {
		return fmt.Sprintf("%s/%s/%s/%s", topicName.domain, topicName.tenant, topicName.namespacePortion, topicName.encodedLocalName())
	} else {
		return fmt.Sprintf("%s/%s/%s/%s/%s", topicName.domain, topicName.tenant, topicName.cluster, topicName.namespacePortion, topicName.encodedLocalName())
	}
}

func (topicName *TopicName) encodedLocalName() string {
	return url.PathEscape(topicName.localName)
}
