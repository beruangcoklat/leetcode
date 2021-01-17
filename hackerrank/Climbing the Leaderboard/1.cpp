#include <bits/stdc++.h>

using namespace std;

vector<string> split_string(string);

struct Node {
  int key; // score
  int rank;
  int height;
  Node *right, *left;
  Node(int key, int rank) {
    this->rank = rank;
    this->key = key;
    height = 1;
    right = left = NULL;
  }
};

int getMax(int a, int b) { return a > b ? a : b; }

int getHeight(Node *x) {
  if (!x)
    return 0;
  return x->height;
}

int getBalance(Node *x) {
  if (!x)
    return 0;
  return getHeight(x->left) - getHeight(x->right);
}

void rotateL(Node *&curr) {
  Node *right = curr->right;
  Node *rightLeft = right->left;

  right->left = curr;
  curr->right = rightLeft;

  curr->height = getMax(getHeight(curr->left), getHeight(curr->right)) + 1;
  right->height = getMax(getHeight(right->left), getHeight(right->right)) + 1;

  curr = right;
}

void rotateR(Node *&curr) {
  Node *left = curr->left;
  Node *leftRight = left->right;

  left->right = curr;
  curr->left = leftRight;

  curr->height = getMax(getHeight(curr->left), getHeight(curr->right)) + 1;
  left->height = getMax(getHeight(left->left), getHeight(left->right)) + 1;

  curr = left;
}

void insertNode(Node *&curr, int key, int rank) {
  if (!curr) {
    curr = new Node(key, rank);
    return;
  } else if (curr->key == key)
    return;
  else if (key > curr->key)
    insertNode(curr->right, key, rank);
  else if (key < curr->key)
    insertNode(curr->left, key, rank);

  curr->height = getMax(getHeight(curr->left), getHeight(curr->right)) + 1;
  int balance = getBalance(curr);
  if (balance > 1 && key < curr->left->key)
    rotateR(curr);
  if (balance > 1 && key > curr->left->key) {
    rotateL(curr->left);
    rotateR(curr);
  }

  if (balance < -1 && key > curr->right->key)
    rotateL(curr);
  if (balance < -1 && key < curr->right->key) {
    rotateR(curr->right);
    rotateL(curr);
  }
}

vector<int> climbingLeaderboard(vector<int> scores, vector<int> alice) {
  Node *root = 0;
  int prevValue;
  int rank = 1;
  for (int s = 0; s < scores.size(); s++) {
    int curr = scores[s];

    if (s == 0) {
      insertNode(root, curr, rank);
      rank++;
      prevValue = curr;
    } else if (curr == prevValue) {
      insertNode(root, curr, rank - 1);
    } else {
      insertNode(root, curr, rank);
      rank++;
      prevValue = curr;
    }
  }

  vector<int> ret;
  for (int i = 0; i < alice.size(); i++) {
    int value = alice[i];
    Node *curr = root;
    int rank;
    while (true) {
      if (value == curr->key) {
        rank = curr->rank;
        break;
      }

      if (value > curr->key && !curr->right) {
        rank = curr->rank;
        break;
      }

      if (value < curr->key && !curr->left) {
        rank = curr->rank + 1;
        break;
      }
      if (value > curr->key) {
        curr = curr->right;
      } else if (value < curr->key) {
        curr = curr->left;
      }
    }
    ret.push_back(rank);
  }

  return ret;
}

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    int scores_count;
    cin >> scores_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string scores_temp_temp;
    getline(cin, scores_temp_temp);

    vector<string> scores_temp = split_string(scores_temp_temp);

    vector<int> scores(scores_count);

    for (int i = 0; i < scores_count; i++) {
        int scores_item = stoi(scores_temp[i]);

        scores[i] = scores_item;
    }

    int alice_count;
    cin >> alice_count;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    string alice_temp_temp;
    getline(cin, alice_temp_temp);

    vector<string> alice_temp = split_string(alice_temp_temp);

    vector<int> alice(alice_count);

    for (int i = 0; i < alice_count; i++) {
        int alice_item = stoi(alice_temp[i]);

        alice[i] = alice_item;
    }

    vector<int> result = climbingLeaderboard(scores, alice);

    for (int i = 0; i < result.size(); i++) {
        fout << result[i];

        if (i != result.size() - 1) {
            fout << "\n";
        }
    }

    fout << "\n";

    fout.close();

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}
