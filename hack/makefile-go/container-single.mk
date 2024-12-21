.PHONY: container
container: ## Build containers
	# COMMIT pass to container, because no git repo in container
	@for target in $(TARGETS); do                                \
	  for registry in $(REGISTRIES); do                          \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX);          \
	    docker build -t $${registry}$${image}:$(VERSION)         \
	      --build-arg COMMIT=$(COMMIT)                           \
	      --progress=plain                                       \
	      -f $(BUILD_DIR)/$${target}/Dockerfile .;               \
	  done                                                       \
	done


.PHONY: push-container
push-container: ## Push containers images to reigstry
	@for target in $(TARGETS); do                       \
	  for registry in $(REGISTRIES); do                 \
	    image=$(IMAGE_PREFIX)$${target}$(IMAGE_SUFFIX); \
	    docker push $${registry}$${image}:$(VERSION);   \
	  done                                              \
	done

