from fastapi import FastAPI
from src.api.routes import register_routes

from src.logger.setup_logger import configure_logger, LogLevels
from src.config.config import Config

config = Config()
print(config.to_dict())

configure_logger(LogLevels.debug)
app = FastAPI()
register_routes(app)
