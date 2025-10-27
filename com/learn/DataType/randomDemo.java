package com.learn.DataType;

import java.util.Random;

public class randomDemo {
    public static void main(String[] args) {
        Random r = new Random();
        int[] arr = new int[10];
        boolean flag = true;

        getRnd:
        for (int i = 0; i < 10; i++) {
            int num = r.nextInt(1, 100) + 1;
            for (int j = 0; j < i; j++) {
                if (num == arr[i]) {
                    flag = false;
                    i--;
                    continue getRnd;
                }
            }
            if (flag) arr[i] = num;
        }

        System.out.print("[");
        for (int n : arr) {
            if (n == arr[arr.length - 1]) System.out.print(n);
            else System.out.print(n + ", ");
        }
        System.out.print("]");
    }
}
