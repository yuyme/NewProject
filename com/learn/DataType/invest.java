package com.learn.DataType;

public class invest {
    public static void a(){
        double principal = 100000; // 初始本金
        double rate = 0.017;       // 年利率 1.7%
        double target = principal * 2; // 目标金额：本金翻倍
        int years = 0;

        while (principal < target) {
            principal = principal * (1 + rate); // 每年计入利息
            years++;
        }

        System.out.println("复利情况下，本金翻倍需要 " + years + " 年。");
    }
    public static void main(String[] args) {
        double new_principal = 100000;
        int count = 0;

        while (new_principal < 200000) {
            new_principal += new_principal * 0.017;
            count ++;
        }

        System.out.println("在第" + count + "年后，将实现本金翻倍!");
        System.out.println("---------------");
        a();


    }
}
