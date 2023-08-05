#! /usr/bin/python3

import inspect
import os
from pprint import pp
from specklepy.api import operations
from specklepy.api.wrapper import StreamWrapper
from specklepy.api.client import SpeckleClient
from specklepy.api.credentials import Account
# from specklepy.api.credentials import get_default_account
from specklepy.objects.base import Base
from specklepy.transports.server import ServerTransport

HOST =  "https://speckle.xyz"
TOKEN = os.environ["SPECKLE_TOKEN"]
PROJECT_NAME = "ExcelTest"
# PROJECT_NAME = "test_webhooks"
# PROJECT_NAME = "ChooseWise"
BUILDING_NAME = "main"
VERSION_ID = "latest"

def main() -> None:
    base = get_base()

    # pp(project.elements[0].__dict__)
    schedule = base["data"]
    headers, data = schedule[0], schedule[1:]

    for d in data:
        pp(d)
    # pp(project[0].["@Walls"])
    # for elem in building.elements[0].elements:
    #     pp(elem["@speckle_type"])
    # pp(inspect.signature(get_project))

    # pp(project.__dict__)
    # pp(project["type"])
    # pp(project.parameters)
    # for slab in project['@IFCSLAB']:
    #     pp(slab)

    # data = get_data(project, "speckle_type")
    # pp(set(data))

def get_base() -> Base:
    """
    Gets the Base object of a Speckle project (stream)
    """
    account = Account() # account = get_default_account()
    account.token = TOKEN

    client = SpeckleClient(HOST)
    client.authenticate_with_account(account)

    project_id = client.stream.search(PROJECT_NAME)[0].id

    transport = ServerTransport(
        client=client,
        stream_id=project_id)

    building = client.branch.get(project_id, BUILDING_NAME)
    if VERSION_ID == "latest":
        version = building.commits.items[0] # latest version
    else:
        version = client.commit.get(project_id, VERSION_ID)

    base =  operations.receive(
        version.referencedObject,
        transport)
    return base


def get_data(base: Base, prop: str) -> list[str]:
    data = [getattr(base, prop)
            if hasattr(base, prop) else None]
    for s in base.elements:
        data.extend(get_data(s, prop))
    return data



if __name__ == "__main__":
    main()
