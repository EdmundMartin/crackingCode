package com.edmundmartin.algos.stringsArrays;

public class Urlify {

    private int countOfChar(char[] str, int start, int end, char target) {
        int count = 0;
        for (int i = start; i < end; i++) {
            if (str[i] == target) {
                count++;
            }
        }
        return count;
    }

    public void replaceSpaces(char[] str, int trueLength) {
        int numberOfSpaces = countOfChar(str, 0, trueLength, ' ');
        int newIndex = trueLength - 1 + numberOfSpaces * 2;

        if (newIndex )
    }
}
