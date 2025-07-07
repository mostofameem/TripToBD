from fastapi.responses import JSONResponse
from typing import Any

def send_error(error_msg: str, error_code: int=500, data: Any = None):
    """Send a standardized error JSON response, with optional data."""
    return JSONResponse(
        status_code=error_code,
        content={
            "message": error_msg,
            "data": data
        }
    )
