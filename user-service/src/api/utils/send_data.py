from fastapi.responses import JSONResponse

def send_data(msg: str, data, status_code: int = 200):
    return JSONResponse(
        status_code=status_code,
        content={
            "message": msg,
            "data": data
        }
    )
