class KataSensor:

    def __init__(self, kata):
        self.kata = kata

daftar_kata = [
    "asu",
    "anjing",
]

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
