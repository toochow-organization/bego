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

k8s_yaml(kustomize('deployments/local')) # load development dependencies such as postgres, etc.

# database
k8s_resource('postgresql', port_forwards='5432:5432', labels=["database"])

# external namespace
include('./external/apis/Tiltfile')