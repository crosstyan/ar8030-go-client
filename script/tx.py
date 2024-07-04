import anyio
import click
import socket
from loguru import logger
from anyio import create_connected_udp_socket, run
from enum import Enum


@click.command()
@click.argument("port", type=int)
def main(port: int):

    async def _main():
        async with await create_connected_udp_socket(family=socket.AF_INET,
                                                     remote_host="localhost",
                                                     remote_port=port) as udp:
            while True:

                def name():
                    if port == 8001:
                        return "DEV"
                    elif port == 9001:
                        return "AP"
                    else:
                        raise ValueError("Invalid port")

                m = f"Hello from {name()}!"
                await udp.send(m.encode())
                logger.info("sent={} remote_port={}", m, port)
                await anyio.sleep(1)

    run(_main)


if __name__ == '__main__':
    main()  # pylint: disable=no-value-for-parameter
