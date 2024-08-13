find . -type f -not -name "rename.sh" -exec perl -pi -e 's|https://github\.com/openshift-online/rh-trex|https://github\.com/eemurphy/brontosaurus|g' {} +
find . -type f -not -name "rename.sh" -not -path "*/.git/*" -exec perl -pi -e 's|rh-trex|brontosaurus|g' {} +
          
# Removing all special characters from the appName so that we can replace all instances of rhtrex and use appName instead
sanitized_name=$(echo "brontosaurus" | tr -cd '[:alnum:]_')
find . -type f -not -name "rename.sh" -not -path "*/.git/*" -exec perl -pi -e 's|rhtrex|'"$sanitized_name"'|ig' {} +

# Replacing trex with appName-service
find . -type f -not -name "rename.sh" -not -path "*/.git/*" -exec perl -pi -e 's|trex|'"brontosaurus-service"'|ig' {} +

# Replacing the old directory names with the app-names 
#mv ./cmd/trex ./cmd/brontosaurus
#mv .tekton/rh-trex-pull-request.yaml ./rh-trex/.tekton/brontosaurus-pull-request.yaml
#mv .tekton/rh-trex-push.yaml ./rh-trex/.tekton/brontosaurus-push.yaml