package com.learn.methodDem;

import java.util.Scanner;

public class Demo1 {

    /*
    *跳水比赛有五个评委打分，分数在0~100之间。最终得分会去掉一个最高分，
     去掉一个最低分，剩余的分数再求平均分，改平均数为选手最终得分。
     * 要求1：利用键盘录入5个整数存入数组中，如果分数超出范围需要重新录入
     * 要求2：定义方法分别求数组最大值和最小值
     * 要求3：计算五名评委的总分
     * 要求4：总分-最大值-最小值，求选手最终平均分
    */

    // 输出分数
    private static void show(int[] arr){
        System.out.print("[");
        for (int i = 0; i < arr.length; i++) {
            if (i == arr.length - 1) System.out.print(arr[i]);
            else System.out.print(arr[i] + ", ");
        }
        System.out.print("]");
    }

    // 拿到所有分数
    private static int[] getScores(Scanner sc){
        // 创建存放五位评委评分的数组
        int[] scores = new int[5];
        // 如果数组存放满五个评分则将sco_count变为true
        int sco_count = 0;
        // 循环获取五次评分 判断每次的分数是否超出范围 然后存放评分数组内
        while (sco_count < 5) {
            // 拿到评分
            System.out.println("请输入第" + (sco_count + 1) + "位评委的评分：");
            int score = sc.nextInt();
            // 判断评分是否超出范围
            if (score < 0 || score > 100) {
                System.err.println("评分超出范围 请重新输入！");
                continue;
            } else {
                scores[sco_count] = score;
                sco_count++;
            }

        }
        return scores;
    }

    // 得到最大值
    private static int getMax(int[] arr){
        // 最高分变量
        int max = arr[0];
        // 遍历分数数组，比最高值大的重新赋值
        for (int a : arr) {
            if (a > max) max = a;
        }
        // 返回最高分
        return max;
    }

    // 得到最小值
    private static int getMin(int[] arr){
        // 最低分变量
        int min = arr[0];
        // 遍历分数数组，比最低值小的重新赋值
        for (int a : arr) {
            if (a < min) min = a;
        }
        // 返回最高分
        return min;
    }

    // 总分
    private static int getSum(int[] arr){
        // 最总分变量
        int sum = 0;
        // 遍历分数数组，比最低值小的重新赋值
        for (int a : arr) {
            sum += a;
        }
        // 返回总分
        return sum;
    }

    // 方法的介绍
    private static void method(int a){
        System.out.println(a);
    }

    private static void method1(int a, int b){
        System.out.println(a + b);
    }

    public static void main(String[] args) {
        // 计算1+2 1+3 5+2 6+4 8+2 的值并打印
        method1(1,2);
        method1(1,3);
        method1(5,2);
        method1(6,4);
        method1(8,2);

        // 结果：
        /*// 创建键盘录取对象
        Scanner sc = new Scanner(System.in);

        //拿到评分数组
        int[] scores = getScores(sc);

        //输出最高分
        int max = getMax(scores);
        System.out.println("最高分为：" + max);

        //输出最低分
        int min = getMin(scores);
        System.out.println("最低分为：" + min);

        //输出总分
        int sum = getSum(scores);
        System.out.println("总分为：" + sum);

        //输出平均分
        int avg = (sum - max + min) / (scores.length - 2);
        System.out.println("平均分为：" + avg);

        // 输出评分
        show(scores);*/

    }
}
