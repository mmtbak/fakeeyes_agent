#!/bin/sh 
export FLASK_APP=src/app:app
export FLASK_ENV=development
# flask run
# cd app && python3 app.py
cur_dir=$(cd "$(dirname "$0")"; pwd)
echo $cur_dir
cd $cur_dir/app && ../myenv/bin/python3 app.py