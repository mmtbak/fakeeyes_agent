
from flask import request
from flask import Flask
from motion import Motion
from protos.command_pb2 import DeviceOperation, OperateCode

import response
app = Flask(__name__)
app.add_url_rule
# router
# app.add_url_rule




@app.route("/api/healthcheck", methods = ["POST", "GET"])
def healthcheck():
    return response.success(None)



@app.route("/api/motion", methods = ["POST"])
def motion():
    try :
        # header  = request.headers
        data = request.stream.read()
        op = DeviceOperation()
        op.ParseFromString(data)
        print(op)
        Motion(op)
        return response.success(None)
    except Exception as e :
        return response.failed(e,str(e))



if __name__ == "__main__":
    initgrio()
    app.run()