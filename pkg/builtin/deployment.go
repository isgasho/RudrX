package builtin

var Deployment = `apiVersion: core.oam.dev/v1alpha2
kind: WorkloadDefinition
metadata:
  name: deployments.apps
  annotations:
    oam.appengine.info/apiVersion: "apps/v1"
    oam.appengine.info/kind: "Deployment"
spec:
  definitionRef:
    name: deployments.apps
  extension:
    template: |
      #Template: {
      	apiVersion: "apps/v1"
      	kind:       "Deployment"
      	metadata: name: deployment.name
      	spec: {
      		containers: [{
      			image: deployment.image
      			name:  deployment.name
      			env:   deployment.env
      			ports: [{
      				containerPort: deployment.port
      				protocol:      "TCP"
      				name:          "default"
      			}]
      		}]
      	}
      }
      deployment: {
      	name:  string
      	image: string
      	port:  *8080 | int
      	env: [...{
      		name:  string
      		value: string
      	}]
      }`
