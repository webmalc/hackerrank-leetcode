#include <assert.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *readline();

int count_char(char *str) {
  long i, count = 0;
  for (i = 0; str[i] != '\0'; ++i) {
    if ('a' == str[i])
      ++count;
  }
  return count;
}

long repeatedString(char *s, long n) {
  int s_len = strlen(s);
  if (s_len > n) {
    s[s_len - (s_len - n)] = 0;
    return count_char(s);
  } else {
    int reminder = n % s_len;
    long count = n / s_len;
    int a_in_string = count_char(s);
    int tail = 0;
    if (reminder != 0) {
      s[s_len - (s_len - reminder)] = 0;
      tail = count_char(s);
    }
    return a_in_string * count + tail;
  }
}

int main() {
  char *s = readline();

  char *n_endptr;
  char *n_str = readline();
  long n = strtol(n_str, &n_endptr, 10);

  if (n_endptr == n_str || *n_endptr != '\0') {
    exit(EXIT_FAILURE);
  }

  long result = repeatedString(s, n);
  printf("\n%ld", result);

  return 0;
}

char *readline() {
  size_t alloc_length = 1024;
  size_t data_length = 0;
  char *data = malloc(alloc_length);

  while (true) {
    char *cursor = data + data_length;
    char *line = fgets(cursor, alloc_length - data_length, stdin);

    if (!line) {
      break;
    }

    data_length += strlen(cursor);

    if (data_length < alloc_length - 1 || data[data_length - 1] == '\n') {
      break;
    }

    size_t new_length = alloc_length << 1;
    data = realloc(data, new_length);

    if (!data) {
      break;
    }

    alloc_length = new_length;
  }

  if (data[data_length - 1] == '\n') {
    data[data_length - 1] = '\0';
  }

  data = realloc(data, data_length);

  return data;
}
