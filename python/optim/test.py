def fuzzy_search_ratio(target_string, sfilter, regex=False):

    tstring = target_string

    # # process regex here. It's a yes no situation (100 or 0)
    # # TODO: HANDLE EXCEPTION
    # if regex:
    #     try:
    #         if re.search(sfilter, tstring):
    #             return 100
    #     except Exception:
    #         pass
    #     return 0

    my_string = "Some"
    input_string = "thisissome"
    ratios = {
        # 100 for identical matches
        "sfilter == tstring": 100,
        # 98 to 99 reserved (2 scores)
        # 97 for identical non-case-sensitive matches
        "sfilter.lower() == tstring.lower()": 97,
        # 95  to 96 reserved (2 scores)
        # 93 to 94 for inclusion matches
        "sfilter in tstring": 94,
        "sfilter.lower() in tstring.lower()": 93,
        # 91  to 92 reserved (2 scores)
        # 80 to 90 for parts matches
        "all(x in tstring.split() for x in sfilter.split())": 90,
        # 88 to 89 reserved (2 scores)
    }
    for r, s in ratios.items():
        if not(eval(r)): continue
        print(s)
        break

    lower_tstring_parts = [x.lower() for x in tstring_parts]
    lower_sfilter_parts = [x.lower() for x in sfilter_parts]
    # exclude override
    if any(x[0] == '!' for x in sfilter_parts):
        exclude_indices = [
            lower_sfilter_parts.index(i) for i in lower_sfilter_parts
            if i[0] == '!'
        ]
        exclude_indices.reverse()
        exclude_list = [
            lower_sfilter_parts.pop(i) for i in exclude_indices
        ]
        for e in exclude_list:
            # doesn't contain
            if len(e) > 1:
                exclude_string = e[1:]
                if any(
                        [exclude_string in
                        part for part in lower_tstring_parts]
                ):
                    return 0
    if all(x in lower_tstring_parts for x in lower_sfilter_parts):
        return 87

    # 85 to 86 reserved (2 scores)

    if all(x in tstring for x in sfilter_parts):
        return 84

    # 82 to 83 reserved (2 scores)

    if all(x in lower_tstring for x in lower_sfilter_parts):
        return 81

    # 80 reserved

    return 0

if __name__ == "__main__":
    fuzzy_search_ratio("thisIsSomestring", "some")
