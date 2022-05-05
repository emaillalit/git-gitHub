# Find all list items
def index_all(search_list, item):
    indices = list()
    for i in range(len(search_list)):
        print('i -> %s' %i)
        if search_list[i] == item:
            print('search_list i -> %s' %search_list[i])
            indices.append([i])
        elif isinstance(search_list[i], list):
            for index in index_all(search_list[i], item):
                print('index -> %s' %index)
                indices.append([i]+index)
    return indices
