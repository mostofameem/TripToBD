from fastapi import APIRouter


router = APIRouter(
    prefix="/auth",
    tags=["Auth"]
)

@router.get("/")
def read_root():
    return {"Hello": "World"}