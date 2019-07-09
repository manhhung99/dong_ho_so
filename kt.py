import time
import os
import datetime

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
# out: string "07:35:43"
def read_system_clock():
   full_time_str = str(datetime.datetime.now())
   hhmmss = full_time_str[11:19]
   return hhmmss

# in: string "07:35:43"
#     font
# out: [ [], [], [] ...]
def convert_characters_to_fonts(time_str, font) :
   out = []
   for i in range(len(time_str)):
       character = time_str[i]
       converted_font = convert_1_character_to_1_font(character, font)
       out.append(converted_font)
   return out


def merge_2_font(font1, font2):
   font = [""] * len(font1)

   for i in range(len(font1)):
       font[i] = font1[i]+font2[i]

   return font

# int:
# out: ["row_1", "row_2", ... ]
def merge_array_of_fonts(arr_of_fonts):
   number_of_row = len(arr_of_fonts[0])

   merged_font = [""]*number_of_row

   for i in range(len(arr_of_fonts)):
       merged_font = merge_2_font(merged_font, arr_of_fonts[i])

   return merged_font


# int: character & font
# out: [] tuong ung voi 1 font
def convert_1_character_to_1_font(character, font):
   if character == "0":
       return font[0]
   if character == "1":
       return font[1]
   if character == "2":
       return font[2]
   if character == "3":
       return font[3]
   if character == "4":
       return font[4]
   if character == "5":
       return font[5]
   if character == "6":
       return font[6]
   if character == "7":
       return font[7]
   if character == "8":
       return font[8]
   if character == "9":
       return font[9]   
   if character == ":":
       return font[10]

def display_arr(arr):
   for i in range(len(arr)):
       row_i = arr[i]
       print(row_i)

def clear_screen():
   os.system("clear")

def sleep(sleep_time):
   time.sleep(sleep_time)

pre_sec=0
while True:
    a= datetime.datetime.now().second
    if pre_sec!= a:
        clear_screen()
        time_str = read_system_clock()
        converted = convert_characters_to_fonts(time_str, font)
        merged = merge_array_of_fonts(converted)
        display_arr(merged)
    pre_sec=a      
    sleep(0.001)