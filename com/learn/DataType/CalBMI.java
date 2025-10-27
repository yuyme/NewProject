package com.learn.DataType;

import java.util.Scanner;

public class CalBMI {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println("请输入体重(kg)：");
        double weight = sc.nextDouble();
        weight =  weight >= 100 ? weight / 2 : weight;

        System.out.println("请输入身高(m)：");
        double height = sc.nextDouble();
        height = height >= 2 ? height / 100 : height;

        double BMI = weight / (height * height);
        String flog ;
        String risk ;

        if (BMI < 18.5){
            flog = "消瘦";
            risk = "部分增加";
        }else if (BMI < 23.9){
            flog = "正常";
            risk = "正常";
        }else if (BMI < 26.9){
            flog = "偏胖";
            risk = "增加";
        }else if (BMI < 29.9){
            flog = "肥胖";
            risk = "中度增加";
        }else {
            flog = "严重肥胖";
            risk = "严重增加";
        }
        System.out.println("您的BMI值为：" + BMI + " 身体状态：" + flog + "健康风险：" + risk);
    }
}
