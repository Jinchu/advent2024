import sys


class Coordinates:
    def __init__(self, x, y):
        self.X = x
        self.Y = y

    def add_padding(self, raw_string):
        padded = ""
        string_len = len(raw_string)
        if string_len == 1:
            padded = "00" + raw_string
        elif string_len == 2:
            padded = "0" + raw_string
        elif string_len == 3:
            padded = raw_string
        else:
            print(f"ERRRO: Cannot padd this: {raw_string}")

        return padded

    def get_str(self):
        raw_x = str(self.X)
        raw_y = str(self.Y)

        string_x = self.add_padding(raw_x)
        string_y = self.add_padding(raw_y)

        return string_x + "," + string_y

    def get_str_w_directions(self, direction):
        only_coord = self.get_str()
        return only_coord + "," + direction


def get_coordinates(input_lines, target):
    y = 0

    coordinates = []
    for current_line in input_lines:
        x = 0
        for letter in current_line:
            if letter == target:
                block = Coordinates(x, y)
                coordinates.append(block)
            x += 1
        y += 1

    return coordinates


def get_grid_size(input_lines):
    x = len(input_lines[0])
    y = len(input_lines)

    return Coordinates(x,y)


def get_input(file_name):
    with open(file_name, 'r', encoding='UTF-8') as handle:
        input_lines = handle.readlines()

    return input_lines

class GuardRoute:
    def __init__(self, map_size, position, direction):
        self.map_size = map_size
        self.position = position
        self.trail = {}
        self.direction = direction

    def travel_north_south(self, lab_map, loop_detection):
        previous_position = Coordinates(self.position.X, self.position.Y)
        y = self.position.Y

        while y < self.map_size.Y:
            if y < 0:
                break

            self.position.Y = y

            if lab_map.get((self.position.X, self.position.Y), False):
                self.position = previous_position
                return 0, self

            if loop_detection:
                trail_point = self.position.get_str_w_directions(self.direction)
                if self.trail.get(trail_point):
                    return 2, self
            else:
                trail_point = self.position.get_str()

            self.trail[trail_point] = True
            previous_position = Coordinates(self.position.X, self.position.Y)

            if self.direction == "north":
                y -= 1
            elif self.direction == "south":
                y += 1
            else:
                raise ValueError("unexpected direction")

        self.position = previous_position
        return 1, self

    def travel_east_west(self, lab_map, loop_detection):
        previous_position = Coordinates(self.position.X, self.position.Y)
        x = self.position.X

        while x < self.map_size.X:
            if x < 0:
                break

            self.position.X = x

            if lab_map.get((self.position.X, self.position.Y), False):
                self.position = previous_position
                return 0, self

            if loop_detection:
                trail_point = self.position.get_str_w_directions(self.direction)
                if self.trail.get(trail_point):
                    return 2, self
            else:
                trail_point = self.position.get_str()

            self.trail[trail_point] = True
            previous_position = Coordinates(self.position.X, self.position.Y)

            if self.direction == "east":
                x += 1
            elif self.direction == "west":
                x -= 1
            else:
                raise ValueError("unexpected direction")

        self.position = previous_position
        return 1, self

    def travel(self, lab_map, loop_detection):
        if self.direction == "north" or self.direction == "south":
            return self.travel_north_south(lab_map, loop_detection)
        elif self.direction == "east" or self.direction == "west":
            return self.travel_east_west(lab_map, loop_detection)
        else:
            raise ValueError("unexpected direction")

    def guard_navigation(self, block_map, loop_detection):
        exit_found = 0
        kill_switch = 1200
        i = j = 0
        all_directions = ["north", "east", "south", "west"]

        while True:
            if j > kill_switch:
                raise RuntimeError("The kill switch was triggered")

            self.direction = all_directions[i]
            exit_found, _ = self.travel(block_map, loop_detection)

            if exit_found == 1:
                return 0
            if exit_found == 2:
                return 1

            i = (i + 1) % len(all_directions)
            j += 1

def main():
    print("Hello")

    input = get_input("./day_06/test-input-1.txt")
    block_coordinates = get_coordinates(input, "#")
    starting_point = get_coordinates(input, "^")
    map_size = get_grid_size(input)

    block_map = {}
    for block in block_coordinates:
        block_str = block.get_str()
        block_map[block_str] = True

    route = GuardRoute(map_size=map_size, position=starting_point[0], direction="north")
    route.guard_navigation(block_map, False)
    return 0


if __name__ == "__main__":
    sys.exit(main())
