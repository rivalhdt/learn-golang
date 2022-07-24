from daftar_kata import daftar_kata
from class_sensor import KataSensor

sensor = KataSensor(daftar_kata)


def sensor_kata(p_phrase):
    result = ""
    if p_phrase in sensor.kata:
        for i in p_phrase:
            if i in "aiueo":
                a = i.replace(i, "*")
                result += a
            else:
                result += i
    return result


is_on = True
while is_on:
    phrase = input("Masukan kata: ")
    print(sensor_kata(phrase))
    if phrase == "tambah":
        tambah = input("Tambahkan kata: ")
        sensor.tambah_kata([tambah])
    elif phrase == "cek":
        print(sensor.kata)
