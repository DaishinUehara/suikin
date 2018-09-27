#!/bin/sh

GOCMD=go

CWD=`pwd`

CMDSRCROOT=${CWD}/cmd
if [ ! -d "${CMDSRCROOT}" ]; then
  echo "folder stracture error" 1>&2
  exit 1
fi

# folder search
cd ${CMDSRCROOT}
CMDDIRS=$(find . -maxdepth 1 -type d -print | sed -e 's/^\.//' | sed -e 's/^\///')
cd ${CWD}

#echo "${CMDDIRS}"

${GOCMD} test -v "${CMDSRCROOT}/skcmnlib/skcmnlib_test.go"

for CMDDIR in ${CMDDIRS}; do
  ${GOCMD} build "${CMDSRCROOT}/${CMDDIR}/${CMDDIR}.go"
done


for CMDDIR in ${CMDDIRS}; do
  ${GOCMD} install "${CMDSRCROOT}/${CMDDIR}/${CMDDIR}.go"
done

