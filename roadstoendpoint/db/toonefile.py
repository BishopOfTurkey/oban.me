import re
import sqlite3

words = open("words.txt")

con = sqlite3.connect("database.db")
cur = con.cursor()

i = 0
for line in words:
    line = line.strip()
    m = re.search("[a-z]+", line)
    if m:
        if line == m.group(0):
            cur.execute("INSERT INTO dictionary(word) VALUES ('"+line+"')")

con.commit()
con.close()
