import anyio
import click
import socket
from loguru import logger
from anyio import create_udp_socket, run
from enum import Enum


@click.command()
@click.argument("port", type=int)
def main(port: int):

    async def _main():
        async with await create_udp_socket(family=socket.AF_INET,
                                           local_port=port) as udp:
            async for packet, (host, remote_port) in udp:
                logger.info("len={}, host={}, port={}, pkt={}", len(packet),
                            host, remote_port, packet.decode())

    run(_main)


if __name__ == '__main__':
    main()  # pylint: disable=no-value-for-parameter
