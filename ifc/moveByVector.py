#! ./env/bin/python3
import ifcopenshell as ifc

def info(obj):
    for k, v in obj.get_info().items():
        print(f"{k}\n\t{v}\n")

def main():
    # TODO file as arg
    file = ifc.open("./rac_advanced_sample_project.ifc")

    doors = file.by_type("ifcdoor")
    for d in doors:
        pset = ifc.util.element.get_psets(d)["Pset_DoorCommon"]
        # p = "AcousticRating"
        p = "FireRating"
        if not p in pset.keys():
            continue
        print(pset[p])
        # print("asdf")

    # walls = file.by_type("ifcwall")
    # seen = set()
    # for w in walls:
    #     if w.ObjectType in seen: continue
    #     seen.add(w.ObjectType)

    # project = file.by_type("ifcproject")[0]
    # info(project)




    # materials = ifc.util.element.get_material(w)
    # for m in materials:
    #     if type(m) in (float, tuple, str):
    #         continue
    #     if m.LayerSetName in seen: continue
    #     seen.add(m.LayerSetName)
    # print(w)
    # print(seen)

if __name__ == "__main__":
    main()

