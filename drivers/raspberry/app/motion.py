
import opcode
from protos import command_pb2
from protos.command_pb2 import DeviceOperation



def Motion(op :DeviceOperation):
    if op.opcode == command_pb2.MoveFront:
        print("front")
    elif op.opcode  == command_pb2.MoveBack:
        print("back")
    elif op.opcode  == command_pb2.TurnRight:
        print("rigth")
    elif op.opcode  == command_pb2.TurnLeft:
        print("left")
    else :
        print("unkonw op : ", op.code )



def MoveForword():
    pass 

def MoveBack():
    pass 

def TurnLeft():
    pass 

def TurnRight():
    pass 