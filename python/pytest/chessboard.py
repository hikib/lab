def getColor(c, r):
    if not 1 <= c <= 8 or not 1 <= r <= 8:
        return ""
    if c % 2 == r % 2:
        return "white"
    else:
        return "black"

def test_getColor():
    assert getColor(2, 1) == "black"
    assert getColor(2, 2) == "white"
    assert getColor(2, 0) == ""

