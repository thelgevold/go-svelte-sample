<script>
  import { interval } from "rxjs";
  import { map, tap, take } from "rxjs/operators";

  let current = 0;

  interval(500)
    .pipe(
      map((v) => v * 100),
      tap((v) => (current = v)),
      take(100)
    )
    .subscribe();
</script>

<h4>RxJs Interval Demo</h4>
{current}
