class Solution {
    public boolean isValidSudoku(char[][] board) {
        int N=9;
        HashSet<Character>[] rows = new HashSet[N];
        HashSet<Character>[] columns = new HashSet[N];
        HashSet<Character>[] boxes = new HashSet[N];

        for(int r=0;r<N;r++){
            rows[r] = new HashSet<Character>();
            columns[r] = new HashSet<Character>();
            boxes[r] = new HashSet<Character>();
        }

        for(int r=0;r<N;r++){
            for(int c=0;c<N;c++){
                char val = board[r][c];
                if(val=='.'){
                    continue;
                }

                if(rows[r].contains(val)){
                    return false;
                }

                rows[r].add(val);

                if(columns[c].contains(val)){
                    return false;
                }

                columns[c].add(val);

                //check the index of box

                int idx = (r/3)*3+c/3;
                if(boxes[idx].contains(val)){
                    return false;
                }

                boxes[idx].add(val);
            }
        }
        return true;
    }
}
