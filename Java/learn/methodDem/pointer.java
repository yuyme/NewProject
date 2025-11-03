package learn.methodDem;

public class pointer {
    /*
     * 给定一个递增有序数组，去除其中重复元素
     *
     * */
    public static void main(String[] args) {
        // 有序数组
        int[] arr = {1, 1, 2, 2, 2, 2, 3, 3, 3, 3};

        int left = 0;
        int right;

        for (int i = 0; i < arr.length; i++) {
            right = i;
            if (arr[right] != arr[left]){
                left++;
                arr[left] = arr[right];
            }
        }

        for (int j : arr) {
            System.out.println(j + " ");
        }


    }


}

