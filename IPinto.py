import IPy


def IP_list(filename):
    IPList = []
    with open(filename, "r") as file:
        for line in file:
            if line != '\n':
                line = line.split('\n')[0].split(' ')
                #line = line.replace(' ', '')
                # line = line.split(',')[2]
                if any(line):  # 去除空元素IP
                    IPList.append(line)
    return IPList


def create_network_segment(IP_list):
    IP_segment_lists = []
    IP_segment_lists_no = []
    for i in IP_list:
        # if len(i) != 2:
        #    print(i)
        IP_make_net = i[0] + "-" + i[1]
        try:
            IP_segment_lists.append(IPy.IP(IP_make_net, make_net=True))
        except Exception as e:
            IP_segment_lists_no.append(i)
            continue
    return IP_segment_lists, IP_segment_lists_no






def int2ip(digit):
    result = []
    for i in range(4):
        digit, mod = divmod(digit, 256)
        result.insert(0, mod)
    return '.'.join(map(str, result))


def create_segment(IP_lists):
    IP_lists2 = [IP_lists[0]]
    for i in range(1, len(IP_lists)):
        IP_int1 = int(IPy.IP(IP_lists2[-1][1]).strDec()) + 1
        IP_int2 = int(IPy.IP(IP_lists[i][0]).strDec())
        if IP_int1 == IP_int2:
            IP_lists2[-1][1] = IP_lists[i][1]
            #print("ture")
        else:
            IP_lists2.append(IP_lists[i])
            #print("false")
    return IP_lists2


if __name__ == "__main__":
    IP_lists = IP_list("./IP.txt")
    
    # print(len(IP_list))
    # print(IP_lists)
    IP_lists2 = create_segment(IP_lists)
    #print(IP_lists2)
    print(len(IP_lists2))
    IP_segment_lists, IP_segment_lists_no = create_network_segment(IP_lists2)
    #segment_into_segment(IP_segment_lists)
    #print([i for i in IP_segment_lists])
    #print(IP_segment_lists_no)
    #print(len(IP_lists2))
    #print(len(IP_segment_lists_no))
    with open('可合并.txt', "w") as fp:
        for i in IP_segment_lists:
            #print(i)
            fp.write(str(i)+"\n")

    with open('行合并.txt', 'w') as f:
        for i in IP_lists2:
            line = str(i[0]) + "    " + str(i[1])
            f.write(line+"\n")

    with open('无法合并.txt', 'w') as fd:
        for i in IP_segment_lists_no:
            line = str(i[0]) + "    " + str(i[1])
            fd.write(line+"\n")
    
    print("OK")
