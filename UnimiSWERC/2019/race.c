#include <stdio.h>
#include <stdlib.h>
#include <math.h>

typedef struct
{
    int id;
    int head_start;
    int speed;
    int current_pos;
} Racer;

int compareRacers(const void *a, const void *b)
{
    const Racer *racerA = (const Racer *)a;
    const Racer *racerB = (const Racer *)b;
    if (racerB->current_pos == racerA->current_pos)
    {
        return racerB->id - racerA->id;
    }
    return racerB->current_pos - racerA->current_pos;
}

int main()
{
    int n, size, head_start, speed;
    scanf("%d", &n);
    Racer *racers = malloc(n * sizeof(Racer));
    for (int i = 0; i < n; i++)
    {
        scanf("%d %d", &head_start, &speed);
        racers[i] = (Racer){i, head_start, speed, head_start + speed};
    }

    size = n;
    for (int i = 0; i < (int)log2(n); i++)
    {
        qsort(racers, size, sizeof(Racer), compareRacers);
        size /= 2;
        for (int i = 0; i < size; i++)
        {
            racers[i].current_pos += racers[i].speed;
        }
    }

    printf("%d\n", racers[0].id);
    return 0;
}