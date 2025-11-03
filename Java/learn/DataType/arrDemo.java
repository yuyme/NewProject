package learn.DataType;

import java.util.Scanner;

public class arrDemo {
    public static void main(String[] args) {
        int[] arr = {33, 5, 22, 44, 55, 33, 47, 34 , 99, 5};


        Scanner sc = new Scanner(System.in);
        System.out.println("请输入任意一个两位整数：");
        int num = sc.nextInt();

        boolean flag = false;
        int index = 0;

        for (int i = 0; i < arr.length; i++) {
            if (num == arr[i]) {
                flag = true;
                index = i;
                break;
            }
        }
        if (flag) System.out.println(index);
        else System.out.println("该数据不存在");
    }
}
