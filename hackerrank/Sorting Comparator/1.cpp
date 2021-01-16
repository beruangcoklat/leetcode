

class Checker{
public:
  	// complete this method
    static int comparator(Player a, Player b)  {
        if(a.score < b.score){
            return -1;
        }

        if(a.score == b.score){
            if(a.name.compare(b.name) > 0){
                return -1;
            }
            else if(a.name.compare(b.name) == 0){
                return 0;
            }
            else{
                return 1;
            }
        }


        return 1;
    }
};

