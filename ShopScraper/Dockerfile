FROM python:3.12-slim

WORKDIR /app

ENV POETRY_VERSION=2.1.2
RUN pip install --no-cache-dir poetry==$POETRY_VERSION

COPY pyproject.toml poetry.lock ./

RUN poetry install --no-interaction --no-ansi --no-root

COPY . .

ENTRYPOINT ["poetry", "run", "bash"]