# create font of characters:
num_0 = [" _ _  ",
    "|   | ",
    "|_ _| "]
num_1 = ["   ",
    "/| ",
    " | "]
num_2 = [" _ _  ",
    " _ _| ",
    "|_ _  "]
num_3 = ["_ _  ",
    "_ _| ",
    "_ _| "]
num_4 = ["      ",
    "|_ _| ",
    "    | "]
num_5 = [" _ _  ",
    "|_ _  ",
    " _ _| "]
num_6 = [" _ _  ",
    "|_ _  ",
    "|_ _| "]
num_7 = ["_ _ ",
    "  / ",
    " /  "]
num_8 = [" _ _  ",
    "|_ _| ",
    "|_ _| "]
num_9 = [" _ _  ",
    "|_ _| ",
    " _ _| "]       
symbol = ["    ",
     " *  ",
     " *  "]

font = [num_0, num_1, num_2, num_3, num_4, num_5, num_6, num_7, num_8, num_9, symbol]
# in: datetime of OS
# out: string "15:30:22"
def read_system_clock()
    full_time_str =Time.now.to_s
    puts full_time_str
    hhmmss=""
    for i in 11..18 do
        hhmmss+=full_time_str[i]
    end
    return hhmmss
end 


#in : string "15:30:22" and font
#out : [[],[],[],...]
def convert_characters_to_fonts(time_str,font)
    out=[]
    for i in 0..(time_str.length-1) do
        character=time_str[i]
        converted_font=convert_1_character_to_1_font(character,font)
        out.push(converted_font)
    end
    return out
end

# int:
# out: ["row_1", "row_2", ... ]
def merge_arr (arr)
    merge=["","","","","","","","","",""]
    for i in 0..(arr[0].length-1) do    
        for j in 0..(arr.length-1) do
            font=arr[j]
            merge[i]+=font[i]
        end
    end
    return merge
end

# int: character & font
# out: [] tuong ung voi 1 font
def convert_1_character_to_1_font(character,font)
    if character=="0"
        return font[0]
    end
    if character=="1"
        return font[1]
    end
    if character=="2"
        return font[2]
    end
    if character=="3"
        return font[3]
    end
    if character=="4"
        return font[4]
    end
    if character=="5"
        return font[5]
    end
    if character=="6"
        return font[6]
    end
    if character=="7"
        return font[7]
    end
    if character=="8"
        return font[8]
    end
    if character=="9"
        return font[9]
    end
    if character==":"
        return font[10]
    end
end
while true do
    system("clear")
    time_str=read_system_clock()
    converted=convert_characters_to_fonts(time_str,font)
    merged=merge_arr(converted)
    puts merged
    sleep(1)
end

