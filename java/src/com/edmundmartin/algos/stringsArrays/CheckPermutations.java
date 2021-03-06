package com.edmundmartin.algos.stringsArrays;

public class CheckPermutations {

    private String sort(String s) {
        char[] content = s.toCharArray();
        java.util.Arrays.sort(content);
        return new String(content);
    }

    public boolean permutationsUsingSorting(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        return sort(s).equals(sort(t));
    }

    public boolean permutationsUsingCharCount(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        int[] letters = new int[128];
        for (int i = 0; i < s.length(); i++) {
            letters[s.charAt(i)]++;
        }

        for (int i = 0; i < t.length(); i++) {
            letters[t.charAt(i)]--;
            if (letters[t.charAt(i)] < 0) {
                return false;
            }
        }
        return true;
    }
}
