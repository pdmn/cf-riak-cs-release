package riak_cs_service

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/vito/cmdtest/matchers"
	"github.com/vito/cmdtest"
	"github.com/pivotal-cf-experimental/cf-test-helpers/runner"
)

var _ = Describe("Riak CS Nodes Register a Route", func() {

		It("Allows users to access the riak-cs service using external url instead of IP of single machine after register the route", func() {
			endpointURL := IntegrationConfig.RiakCsScheme + "riakcs." + IntegrationConfig.AppsDomain + "/riak-cs/ping"

			var session *cmdtest.Session
			session = runner.Curl("-k", endpointURL)

			Expect(session).To(SayWithTimeout("OK", 60*time.Second))
		})
	})

var _ = Describe("Riak Broker Registers a Route", func() {

		It("Allows users to access the riak-cs broker using a url", func() {
				endpointURL := "http://p-cs." + IntegrationConfig.SystemDomain + "/v2/catalog"

				var session *cmdtest.Session
				session = runner.Curl("-k", "-s", "-w", "%{http_code}", endpointURL, "-o", "/dev/null")

				// check for 401 because it means we reached the endpoint, but did not supply credentials.
				// a failure would be a 404
				Expect(session).To(SayWithTimeout("401", 60*time.Second))
			})
	})
