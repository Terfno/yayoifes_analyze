#!/bin/bash

make run

if [ "$?"=1 ]; then
  ./run.sh
fi
