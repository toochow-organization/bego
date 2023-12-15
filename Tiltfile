## Tiltfile configure a deployment of our application in a local kubernetes
allow_k8s_contexts([
  'docker-desktop',
  'rancher-desktop'
])

# Overriding the default kustomize function to allow
# passing kustomize build options.
# Default builtin kustomize do not allow passing options.
def kustomize(path):
  cmd = "kustomize build --enable-helm " + path
  return local(cmd, command_bat=cmd, quiet=True)

# Run kustomize on a given directory and return the resulting YAML as a Blob Directory is watched.
# We pipe the output of kustomize into k8s_yaml directive and deploy the YAML to the local Kubernetes.
k8s_yaml(kustomize('deployments/local')) # load development dependencies such as postgres, gcp emulators etc.

# database
k8s_resource('postgresql', port_forwards='5432:5432', labels=["database"])

