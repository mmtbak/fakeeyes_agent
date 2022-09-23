

from typing import Any


def success(data:Any ) ->dict:
    return {
        "Success": True, 
        "Data": data ,
        "ErrorCode": None, 
        "ErrorMessage": None
    }

def failed(code: str, message: str ) -> dict:
    return {
        "Success": False,
        "Data": None ,
        "ErrorCode": code, 
        "ErrorMessage": message
    }