## Model

```plantuml
-> module ++: OfFuncT(context, func() T)
module -> CallService **: create
return chanT
box Thread
    participant CallService
end box
loop until context ends
CallService -> : T = func()
ref over CallService
chanT <- T
end
end

```
