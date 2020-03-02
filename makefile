#make test

OUTPUT_ROOT := output
MKDIR_P = mkdir -p
RM = rm -rf
CP = cp -r
BLOOMFLITER = bloomfilter
ROUTER = router

bf:
	@echo "build bloomfilter binary ....."
	${MKDIR_P} ${OUTPUT_ROOT}/${BLOOMFLITER}
	cd cmd/${BLOOMFLITER};go build -o ../../${OUTPUT_ROOT}/${BLOOMFLITER}/${BLOOMFLITER} -v .
	${CP} config ${OUTPUT_ROOT}/${BLOOMFLITER}
	@echo "build finish"

router:
	@echo "build router binary ....."
	${MKDIR_P} ${OUTPUT_ROOT}/${ROUTER}
	cd cmd/${ROUTER};go build -o ../../${OUTPUT_ROOT}/${ROUTER}/${ROUTER} -v .
	@echo "build finish"
	${CP} config ${OUTPUT_ROOT}/${ROUTER}

clean:
	${RM} ${OUTPUT_ROOT}