from app import run_app
import argparse


if __name__ == '__main__':
    parser = argparse.ArgumentParser()

    parser.add_argument('-d', '--debug', action='store_true', dest='debug', help='set server into debug mode')
    parser.add_argument('-p', '--port', type=int, dest='port', help='set port')

    args = parser.parse_args()

    run_app(args.__dict__)
