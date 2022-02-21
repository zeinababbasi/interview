import unittest
from DetectNumbers.PythonVersion import detect_numbers


class DetectNumberTest(unittest.TestCase):
    def test_numbers(self):
        self.assertIsNotNone(detect_numbers.string_is_number("10"))
        self.assertIsNotNone(detect_numbers.string_is_number("-10"))
        self.assertIsNotNone(detect_numbers.string_is_number("10.1"))
        self.assertIsNotNone(detect_numbers.string_is_number("-10.1"))
        self.assertIsNotNone(detect_numbers.string_is_number("1e5"))

    def test_non_numbers(self):
        self.assertIsNone(detect_numbers.string_is_number("a"))
        self.assertIsNone(detect_numbers.string_is_number("x 1"))
        self.assertIsNone(detect_numbers.string_is_number("a -2"))
        self.assertIsNone(detect_numbers.string_is_number("-"))


if __name__ == '__main__':
    unittest.main()
