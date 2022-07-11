parse_string = ""

group1 = []
group2 = []
group3 = []
group4 = []
group5 = []
group6 = []
group7 = []
group8 = []

result_group1 = []
result_group2 = []
result_group3 = []
result_group4 = []
result_group5 = []
result_group6 = []
result_group7 = []
result_group8 = []

with open('emojihex2.txt') as f:
    lines = [line.rstrip() for line in f]

for l in lines:
    splits = l.split(" ")
    if len(splits) == 1:
        group1.append(splits[0])
    if len(splits) == 2:
        group2.append(splits)
    if len(splits) == 3:
        group3.append(splits)
    if len(splits) == 4:
        group4.append(splits)
    if len(splits) == 5:
        group5.append(splits)
    if len(splits) == 6:
        group6.append(splits)
    if len(splits) == 7:
        group7.append(splits)
    if len(splits) == 8:
        group8.append(splits)

group1.sort()

#group1
for i in range(len(group1)):
    if len(result_group1) == 0:
        result_group1.append(group1[i])
    else:
        if result_group1[len(result_group1)-1] != group1[i]:
            splits = result_group1[len(result_group1)-1].split("-")
            n = int(splits[0], 16)
            if len(splits) == 2:
                n = int(splits[1], 16)
            m = int(group1[i], 16)
            if (n+1) == m:
                result_group1[len(result_group1)-1] = splits[0]+"-"+group1[i]
            else:
                result_group1.append(group1[i])


#group2
for i in range(len(group2)):
    if len(result_group2) == 0:
        result_group2.append(group2[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group2[i])):
            if result_group2[len(result_group2)-1][j] == group2[i][j]:
                result.append(group2[i][j])
            else:
                splits = result_group2[len(result_group2)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group2[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group2[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group2.append(group2[i])
        else:
            result_group2[len(result_group2)-1] = result

#group3
for i in range(len(group3)):
    if len(result_group3) == 0:
        result_group3.append(group3[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group3[i])):
            if result_group3[len(result_group3)-1][j] == group3[i][j]:
                result.append(group3[i][j])
            else:
                splits = result_group3[len(result_group3)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group3[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group3[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group3.append(group3[i])
        else:
            result_group3[len(result_group3)-1] = result


#group4
for i in range(len(group4)):
    if len(result_group4) == 0:
        result_group4.append(group4[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group4[i])):
            if result_group4[len(result_group4)-1][j] == group4[i][j]:
                result.append(group4[i][j])
            else:
                splits = result_group4[len(result_group4)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group4[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group4[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group4.append(group4[i])
        else:
            result_group4[len(result_group4)-1] = result

#group5
for i in range(len(group5)):
    if len(result_group5) == 0:
        result_group5.append(group5[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group5[i])):
            if result_group5[len(result_group5)-1][j] == group5[i][j]:
                result.append(group5[i][j])
            else:
                splits = result_group5[len(result_group5)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group5[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group5[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group5.append(group5[i])
        else:
            result_group5[len(result_group5)-1] = result


#group6
for i in range(len(group6)):
    if len(result_group6) == 0:
        result_group6.append(group6[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group6[i])):
            if result_group6[len(result_group6)-1][j] == group6[i][j]:
                result.append(group6[i][j])
            else:
                splits = result_group6[len(result_group6)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group6[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group6[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group6.append(group6[i])
        else:
            result_group6[len(result_group6)-1] = result

#group7
for i in range(len(group7)):
    if len(result_group7) == 0:
        result_group7.append(group7[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group7[i])):
            if result_group7[len(result_group7)-1][j] == group7[i][j]:
                result.append(group7[i][j])
            else:
                splits = result_group7[len(result_group7)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group7[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group7[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group7.append(group7[i])
        else:
            result_group7[len(result_group7)-1] = result


#group8
for i in range(len(group8)):
    if len(result_group8) == 0:
        result_group8.append(group8[i])
    else:
        result = []
        isDiffrent = False
        for j in range(len(group8[i])):
            if result_group8[len(result_group8)-1][j] == group8[i][j]:
                result.append(group8[i][j])
            else:
                splits = result_group8[len(result_group8)-1][j].split("-")
                n = int(splits[0], 16)
                if len(splits) == 2:
                    n = int(splits[1], 16)
                m = int(group8[i][j], 16)
                if (n+1) == m:
                    result.append(splits[0]+"-"+group8[i][j])
                else:
                    isDiffrent = True
        if isDiffrent:
            result_group8.append(group8[i])
        else:
            result_group8[len(result_group8)-1] = result

#print(result_group1)
# print(result_group2)
# print(result_group3)
# print(result_group4)



#(?:\x{1F468}[\x{1F3FB}-\x{1F3FF}]\x{200D}\x{1F9B0})
for rg5 in result_group8:
    for i in range(len(rg5)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
        else:
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg5[i]+"}"
    parse_string = parse_string + ")|"




for rg5 in result_group7:
    for i in range(len(rg5)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
        else:
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg5[i]+"}"
    parse_string = parse_string + ")|"






for rg5 in result_group6:
    for i in range(len(rg5)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
        else:
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg5[i]+"}"
    parse_string = parse_string + ")|"




for rg5 in result_group5:
    for i in range(len(rg5)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg5[i]+"}"
        else:
            if "-" in rg5[i]:
                splits = rg5[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg5[i]+"}"
    parse_string = parse_string + ")|"




for rg4 in result_group4:
    for i in range(len(rg4)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg4[i]+"}"
            if "-" in rg4[i]:
                splits = rg4[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg4[i]+"}"
        else:
            if "-" in rg4[i]:
                splits = rg4[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg4[i]+"}"
    parse_string = parse_string + ")|"


for rg3 in result_group3:
    for i in range(len(rg3)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg3[i]+"}"
            if "-" in rg3[i]:
                splits = rg3[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg3[i]+"}"
        else:
            if "-" in rg3[i]:
                splits = rg3[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg3[i]+"}"
    parse_string = parse_string + ")|"



for rg2 in result_group2:
    for i in range(len(rg2)):
        if i == 0:
            #parse_string = parse_string + "(?:\\x{"+rg2[i]+"}"
            if "-" in rg2[i]:
                splits = rg2[i].split("-")
                parse_string = parse_string + "(?:[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "(?:\\x{"+rg2[i]+"}"
        else:
            if "-" in rg2[i]:
                splits = rg2[i].split("-")
                parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]"
            else:
                parse_string = parse_string + "\\x{"+rg2[i]+"}"
    parse_string = parse_string + ")|"


for rg1 in result_group1:
    if "-" in rg1:
        splits = rg1.split("-")
        parse_string = parse_string + "[\\x{" + splits[0] +"}-\\x{"+splits[1]+"}]|"
    else:
        parse_string = parse_string + "\\x{" + rg1 +"}|"

f = open("emojiregex.txt", "w")
f.write(parse_string[:-1])
f.close()

print("done")


