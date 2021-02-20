#!/bin/bash -e

dialects=("mysql")

if [[ $(pwd) == *"gorm/tests"* ]]; then
  cd ..
fi

if [ -d tests ]
then
  cd tests
  cp go.mod go.mod.bak
  sed '/^[[:blank:]]*gorm.io\/driver/d' go.mod.bak > go.mod
  cd ..
fi

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    echo "testing ${dialect}..."

    if [ "$GORM_VERBOSE" = "" ]
    then
      if [ -d tests ]
      then
        cd tests
        # with customized tracer
        WITH_TRACER2=true GORM_DIALECT=${dialect} go test -count=1 ./...
        # with global tracer
        GORM_DIALECT=${dialect} go test -race -count=1 ./...
        cd ..
      fi
    fi
  fi
done

if [ -d tests ]
then
  cd tests
  mv go.mod.bak go.mod
fi