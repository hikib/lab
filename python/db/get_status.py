#!/usr/bin python3
import psycopg2

rows = None
status = "status"

try:
    con = psycopg2.connect(
        database="someapp",
        user="hikib",
        password="somepw",
        port="5432",
        host="localhost",
    )

    with con:
        with con.cursor() as cur:
            cur.execute("SELECT * FROM status")
            rows = cur.fetchall()

    # CON CAN BE USED MORE THAN ONCE!
    # with con:
    #     with con.cursor() as cur:
    #         cur.execute("SELECT * FROM status;")
    #         rows = cur.fetchall()

except Exception as e:
    print("conection could not be made due to the following error: \n", e)

finally:
    con.close()

if rows:
    for row in rows:
        print(row)
