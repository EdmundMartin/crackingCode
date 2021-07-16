package com.edmundmartin.algos.stringsArrays;

public class OneAway {

    public boolean oneEditAway(String first, String second) {
        if (first.length() == second.length()) {
            return oneEditReplace(first, second);
        } else if (first.length() + 1 == second.length()) {
            return oneEditInsert(first, second);
        } else if (first.length() - 1 == second.length()) {
            return oneEditInsert(second, first);
        }
        return false;
    }

    private boolean oneEditReplace(String first, String second) {
        boolean foundDifference = false;
        for (int i = 0; i < first.length(); i++) {
            if (first.charAt(i) != second.charAt(i)) {
                if (foundDifference) {
                    return false;
                }
                foundDifference = true;
            }
        }
        return true;
    }

    private boolean oneEditInsert(String first, String second) {
        int firstIndex = 0;
        int secondIndex = 0;
        while (secondIndex < second.length() && firstIndex < first.length()) {
            if (first.charAt(firstIndex) != second.charAt(secondIndex)) {
                if (firstIndex != secondIndex) {
                    return false;
                }
                secondIndex++;
            } else {
                firstIndex++;
                secondIndex++;
            }
        }
        return true;
    }
}
