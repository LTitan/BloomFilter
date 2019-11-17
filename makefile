#make test

OUTPUT_ROOT := output
MKDIR_P = mkdir -p
RM = rm -rf
SLAVE = slave
MASTER = master

slave:
	@echo "build slave ....."
	${MKDIR_P} ${OUTPUT_ROOT}/${SLAVE}
	cd cmd/${SLAVE};go build -o ../../${OUTPUT_ROOT}/${SLAVE}/${SLAVE} -v .
	@echo "build finish"

master:
	@echo "build master ....."
	${MKDIR_P} ${OUTPUT_ROOT}/${MASTER}
	cd cmd/${MASTER};go build -o ../../${OUTPUT_ROOT}/${MASTER}/${MASTER} -v .
	@echo "build finish"

clean:
	${RM} ${OUTPUT_ROOT}