import argparse
import whirlwind_org.server as server


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument(
        '-d', '--debug',
        help='Run website in debug mode',
        dest='debug',
        action='store_true')
    parser.add_argument(
        '-p', '--port',
        help='Specify port to run website on',
        dest='port',
        type=int
    )
    args = parser.parse_args()
    server.run_server(**args.__dict__)

