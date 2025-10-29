package com.learn.DataType;

import java.util.Scanner;

public class invest {
    public static void a(){
        double principal = 100000; // 初始本金
        double rate = 0.017;       // 年利率 1.7%
        double target = principal * 2; // 目标金额：本金翻倍
        int years = 0;

        while (principal < target) {
            principal = principal * (1 + rate); // 每年复利
            years++;
        }

        System.out.println("在这种利率下，本金翻倍需要 " + years + " 年。");
    }

    public static void b(){
        Scanner in = new Scanner(System.in);
        // 获取要计算翻倍的本金
        System.out.println("请输入要计算翻倍的本金：");
        int money = in.nextInt();
        int result = money * 2;

        double pay = money;
        double rate = 0.017;
        int year = 0;

        while (pay < result) {
            pay = pay * (1 + rate);
            year++;
        }
        System.out.println("需要 " + year + " 年");
        in.close();
    }
    public static void main(String[] args) {
        /*double new_principal = 100000;
        int count = 0;

        while (new_principal < 200000) {
            new_principal += new_principal * 0.017;
            count ++;
        }

        System.out.println("在第" + count + "年后，将实现本金翻倍!");
        System.out.println("---------------");
        a();*/

        b();

    }
}
